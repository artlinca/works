package apicupcfg

import (
	"flag"
	"os"
)

const outdir = "outdir"
const managementOut = "management"
const gatewayOut = "gateway"
const analyticsOut = "analytics"
const portalOut = "portal"
const userFacingPublicCertsOut = "userfacingpubliccerts"
const publicCertsOut = "publiccerts"
const mutualAuthCertsOut = "mutualauthcerts"
const commonCertsOut = "commoncerts"
const certbotUserFacingPublicCertOut = "certbotuserfacingpubliccerts"
const certbotPublicCertOut = "certbotpubliccerts"
const etUserFacingPublicCertsOut = "shared-endpoint-trust-userfacingpubliccerts"
const etPublicCertsOut = "shared-endpoint-trust-publiccerts"
const etMutualAuthCertsOut = "shared-endpoint-trust-mutualauthcerts"

const CommonCsrOutDir = "common-csr"
const CustomCsrOutDir = "custom-csr"
const SharedCsrOutDir = "shared-csr"
const ProjectOutDir = "project"

const DatapowerOutDir = "datapower"

func OutputFiles(baseout string) map[string]string {

	outfiles := map[string]string{
		outdir:                         baseout,
		managementOut:                  "apicup-subsys-set-management",
		gatewayOut:                     "apicup-subsys-set-gateway",
		analyticsOut:                   "apicup-subsys-set-analytics",
		portalOut:                      "apicup-subsys-set-portal",
		userFacingPublicCertsOut:       "apicup-certs-set-user-facing-public",
		publicCertsOut:                 "apicup-certs-set-public",
		mutualAuthCertsOut:             "apicup-certs-set-mutual-auth",
		commonCertsOut:                 "apicup-certs-set-common",
		CommonCsrOutDir:                CommonCsrOutDir,
		CustomCsrOutDir:                CustomCsrOutDir,
		SharedCsrOutDir:				SharedCsrOutDir,
		certbotUserFacingPublicCertOut: "apicup-certs-set-certbot-user-facing-public",
		certbotPublicCertOut:           "apicup-certs-set-certbot-public",
		etUserFacingPublicCertsOut:		"apicup-certs-set-shared-trust-user-facing-public",
		etPublicCertsOut:				"apicup-certs-set-shared-trust-public",
		etMutualAuthCertsOut:			"apicup-certs-set-shared-trust-mutual-auth",
		DatapowerOutDir:				DatapowerOutDir,
	}

	return outfiles
}

func tagOutputFileName(outName string, tag string) string {
	const dot = "."
	return outName + dot + tag
}

func concatSubdir(dir1 string, dir2 string) string {
	return dir1 + string(os.PathSeparator) + dir2
}

func Input() (input string, outdir1 string, validateIp bool, initConfig bool, initConfigType string,
	subsysOnly bool, certsOnly bool, certcopy string, certdir string,
	certverify bool, certfile, cafile string, rootcafile string, noexpire bool, certconcat bool, gen bool,
	soma bool, req string, auth string, url string, setfile string, dpdir string, dpfile string,
	datapowerOnly bool) {

	// define command line flags
	inputArg := flag.String("config", "subsys-config.json", "-config input-file")
	outdirArg := flag.String("out", "output", "-out output-directory")
	//commonCsrSubdirArg := flag.String("commoncsr", "common-csr", "-commoncsr subdir")
	//customCsrSubdirArg := flag.String("customcsr", "custom-csr", "-customcsr subdir")
	//projectSubdirArg := flag.String("project", "project", "-project subdir")

	validateIpArg := flag.Bool("validateip", false, "-validateip [true] validate ip addresses")

	initConfigArg := flag.Bool("initconfig", false, "-initconfig [true] initalize json config, use with configtype option")
	initConfigTypeArg := flag.String("configtype", "ova", "-configtype ova|k8s use with initconfig option")

	subsysOnlyArg := flag.Bool("subsys", false, "-subsys [true] generate subsystem scripts only")
	certsOnlyArg := flag.Bool("certs", false, "-certs [true] generate certs scripts only")

	certCopyArg := flag.String("certcopy", "", "-certcopy certfile copy certificate to destination")
	certDirArg := flag.String("certdir", "","-certdir dir copy all certificate files in dir to destination")

	// -certverify [-cert] ... -ca ... -rootca ... -noexpire
	// -certconcat -ca ... -rootca ... -noexpire
	// -gen

	certverifyArg := flag.Bool("certverify", false, "-certverify, verify cert")
	certfileArg := flag.String("cert", "", "-cert cerfile, cert file name")
	cafileArg := flag.String("ca", "", "-ca cafile, ca file, use with -certverify")
	rootcafileArg := flag.String("rootca", "", "-rootca file, root ca file, use with -certverify")
	noexpireArg := flag.Bool("noexpire", false,"-noexpire, check for cert expiration")

	certconcatArg := flag.Bool("certconcat", false, "-certconcat, concatinate ca certs")

	genArg := flag.Bool("gen", false, "-gen, generate scripts")

	somaArg := flag.Bool("soma", false, "-soma, datapower soma request")
	reqArg := flag.String("req", "", "-req somareq.xml, soma request file")
	authArg := flag.String("auth", "dp.env", "-auth envfile, datapower auth file with user and password")
	urlArg := flag.String("url","", "-url datapower soma url")
	setfileArg := flag.String("setfile", "", "-setfile filename, filename to upload to datapower")
	dpdirArg := flag.String("dpdir", "", "-dpdir cert|local|..., datapower directory")
	dpfileArg := flag.String("dpfile","","-dpfile datapower file name")

	datapowerOnlyArg := flag.Bool("datapower", false, "-datapower true|false generate datapower configuration only")

	// scan command line args
	flag.Parse()

	input = *inputArg
	outdir1 = *outdirArg
	//commonCsrSubdir = commonCsrOutDir
	//customCsrSubdir = customCsrOutDir
	//projectSubdir = projectOutDir
	validateIp = *validateIpArg
	initConfig = *initConfigArg
	initConfigType = *initConfigTypeArg
	subsysOnly = *subsysOnlyArg
	certsOnly = *certsOnlyArg
	certcopy = *certCopyArg
	certdir = *certDirArg

	certverify = *certverifyArg
	certfile = *certfileArg
	cafile = *cafileArg
	rootcafile = *rootcafileArg
	noexpire = *noexpireArg
	certconcat = *certconcatArg
	gen = *genArg

	soma = *somaArg
	req = *reqArg
	auth = *authArg
	url = *urlArg
	setfile = *setfileArg
	dpdir = *dpdirArg
	dpfile = *dpfileArg

	datapowerOnly = *datapowerOnlyArg

	return input, outdir1, validateIp, initConfig, initConfigType, subsysOnly, certsOnly, certcopy, certdir,
		certverify, certfile, cafile, rootcafile, noexpire, certconcat, gen,
		soma, req, auth, url, setfile, dpdir, dpfile, datapowerOnly
}
