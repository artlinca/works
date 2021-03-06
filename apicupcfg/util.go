package apicupcfg

import (
	"encoding/json"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/Masterminds/sprig"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func bytesInBox(file string, tbox *rice.Box) []byte {
	if tbox == nil {
		log.Fatal("uninitialized template box")
	}

	c, err := tbox.Bytes(file)

	if err != nil {
		log.Fatal(err)
	}

	return c
}

func FileName(dir string, file string) string {
	return fileName(dir, file)
}

func fileName(dir string, file string) string {
	dir2 := dir

	if len(dir2) == 0 {
		dir2 = "."
	}

	return dir2 + string(os.PathSeparator) + file
}

func isFileExist(file string) (bool, error) {
	_, err := os.Stat(file)

	if os.IsExist(err)  {
		return true, nil

	} else if os.IsNotExist(err) {
		return false, nil
	}

	// pretend file exists, return error
	return true, err
}

func openFile(file string) (*os.File, error) {
	return os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
}

func openFileAppend(file string) (*os.File, error) {
	return os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
}

func readFileBytes(file string) []byte {
	c, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return c
}

func writeFileBytes(file string, bytes []byte) {
	err := ioutil.WriteFile(file, bytes, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateOutputDirectories(outdir, commonCsrSubdir, customCsrSubdir, sharedCsrSubdir, projectSubdir, datapowerSubdir string) error {

	basedir, err := filepath.Abs(outdir)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(basedir, os.ModePerm); err != nil {
		return err
	}

	commonCsrDir := concatSubdir(basedir, commonCsrSubdir)
	if err = os.MkdirAll(commonCsrDir, os.ModePerm); err != nil {
		return err
	}

	customCsrDir := concatSubdir(basedir, customCsrSubdir)
	if err = os.MkdirAll(customCsrDir, os.ModePerm); err != nil {
		return err
	}

	//sharedCsrDir := concatSubdir(basedir, sharedCsrSubdir)
	//if err = os.MkdirAll(sharedCsrDir, os.ModePerm); err != nil {
	//	return err
	//}

	projectDir := concatSubdir(basedir, projectSubdir)
	if err = os.MkdirAll(projectDir, os.ModePerm); err != nil {
		return err
	}

	// bin subdirectory
	binDir := concatSubdir(basedir, "bin")
	if err = os.MkdirAll(binDir, os.ModePerm); err != nil {
		return err
	}

	// datapower subdirectory
	if len(datapowerSubdir) > 0 {
		datapowerDir := concatSubdir(basedir, datapowerSubdir)
		if err = os.MkdirAll(datapowerDir, os.ModePerm); err != nil {
			return err
		}

		// soma subdirectory (this is internal)
		const somasubdir = "soma"
		somadir := concatSubdir(datapowerDir, somasubdir)
		if err = os.MkdirAll(somadir, os.ModePerm); err != nil {
			return err
		}
	}

	return err
}

func tpdir(tbox *rice.Box) string {
	if tbox != nil {
		return ""
	}
	return "templates" + string(os.PathSeparator)
}

func readBytes(file string, tbox *rice.Box) []byte {
	if tbox != nil {
		return bytesInBox(file, tbox)
	}
	return readFileBytes(file)
}

func copyFile(srcfile string, dstfile string) {
	bytes := readFileBytes(srcfile)
	writeFileBytes(dstfile, bytes)
}

func concatFiles2(srcfiles []string, dstfile string) {

	bytes := make([]byte, 0)

	for _, srcfile := range srcfiles {
		bytes1 := readFileBytes(srcfile)

		for _, b := range bytes1 {
			bytes = append(bytes, b)
		}

		// add new-line
		bytes = append(bytes, byte('\n'))
	}

	writeFileBytes(dstfile, bytes)
}

func concatFiles(srcfile1 string, srcfile2 string, dstfile string) {

	bytes1 := readFileBytes(srcfile1)
	bytes2 := readFileBytes(srcfile2)

	bytes := make([]byte, 0, len(bytes1) + len(bytes2))

	for _, b := range bytes1 {
		bytes = append(bytes, b)
	}

	for _, b := range bytes2 {
		bytes = append(bytes, b)
	}

	writeFileBytes(dstfile, bytes)
}

func parseTemplate(tbox *rice.Box, file string) *template.Template {

	str := string(readBytes(file, tbox))
	t := template.New(filepath.Base(file)).Funcs(sprig.TxtFuncMap())

	t, err := t.Parse(str)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func parseTemplates(tbox *rice.Box, file string, files ...string)  *template.Template {

	s := string(readBytes(file, tbox))
	t := template.New(filepath.Base(file)).Funcs(sprig.TxtFuncMap())

	t, err := t.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		s := string(readBytes(f, tbox))

		n := filepath.Base(f)
		t1 := t.New(n).Funcs(sprig.TxtFuncMap())

		t1, err := t1.Parse(s)
		if err != nil {
			log.Fatal(err)
		}
	}
	return t
}

// can do better...
func copySlices(dst []string, src []string) []string {

	if dst == nil && src != nil {
		return src
	}

	if src == nil {
		if dst == nil {
			return []string{}
		} else {
			return dst
		}

	} else {
		for _, v := range src {
			// duplicates possible
			dst = append(dst, v)
		}
	}

	return dst
}

func executeTemplate(t *template.Template, obj interface{}) []string {

	b := executeTemplate2(t, obj)
	return strings.Split(b, "\n")
}

func executeTemplate2(t *template.Template, obj interface{}) string {
	b := &strings.Builder{}

	err := t.Execute(b, obj)
	if err != nil {
		log.Fatal(err)
	}

	return b.String()
}

func writeLines(lines []string, file string) {
	f, err := openFile(file)

	if err == nil {
		defer func() {_ = f.Close()}()
	} else {
		f = os.Stdout
	}

	spacecount := 0
	firstline := true

	yaml := strings.HasSuffix(strings.ToLower(file),".yml" ) ||
		strings.HasSuffix(strings.ToLower(file), "yaml") ||
		strings.HasSuffix(strings.ToLower(file), "json")

	for _, line := range lines {

		if yaml {
			// output yaml file as is...
			// matching on the file extension will not work if output to stdout...
			// look for the yaml file marker at the first line (todo)
			_, _ = fmt.Fprintf(f, "%s\n", line)
			continue
		}

		s := strings.TrimSpace(line)

		if len(s) == 0 {
			spacecount++

		} else if spacecount > 0 {
			// collapse blank lines
			spacecount = 0

			if firstline {
				// no blank line before the fist line
				firstline = false
				_, _ = fmt.Fprintf(f, "%s\n", s)

			} else {
				// insert blank line in front of  aline
				_, _ = fmt.Fprintf(f, "\n%s\n", s)
			}

		} else {
			_, _ = fmt.Fprintf(f, "%s\n", s)

			if firstline {
				firstline = false
			}
		}
	}
}

func writeTemplate(t *template.Template, file string, obj interface{}) {
	lines := executeTemplate(t, obj)
	writeLines(lines, file)
}

func writeTemplate2(t *template.Template, file string, xf func(buf string) string, obj interface{}) {
	buf := executeTemplate2(t, obj)
	writeFileBytes(file, []byte(xf(buf)))
}

func unmarshalJsonFile(file string, objptr interface{}) {
	if err := json.Unmarshal(readFileBytes(file), objptr); err != nil {
		log.Fatalf("json unmarshal failed... %s", err)
	}
}