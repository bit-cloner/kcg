package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"path/filepath"

	"gopkg.in/AlecAivazis/survey.v1"
	//apiv1 "k8s.io/api/core/v1"
	//rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Forming questions

// the questions to ask
var qs = []*survey.Question{

	{
		Name:   "clustername",
		Prompt: &survey.Input{Message: "Enter name of cluster", Help: "This will be used as current context in kubeconfig file"},
	},
}

func main() {
	fmt.Println("\n\t " +
		`888    d8P   .d8888b.   .d8888b.  
         888   d8P   d88P  Y88b d88P  Y88b 
         888  d8P    888    888 888    888 
         888d88K     888        888        
         8888888b    888        888  88888 
         888  Y88b   888    888 888    888 
         888   Y88b  Y88b  d88P Y88b  d88P 
         888    Y88b  "Y8888P"   "Y8888P88 `)
	fmt.Println("\t " +
		`                                          _                                          
 |      |_   _  ._ ._   _ _|_  _   _    _  _  ._ _|_ o  _     _   _  ._   _  ._ _. _|_  _  ._ 
 |< |_| |_) (/_ |  | | (/_ |_ (/_ _>   (_ (_) | | |  | (_|   (_| (/_ | | (/_ | (_|  |_ (_) |  
                                                        _|    _|                              ` + "\n")

	// the answers will be written to this struct
	answers := struct {
		Clustername string
	}{}

	// ask questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// user input is stored in these variables
	var namespace string
	var clustername string
	clustername = answers.Clustername
	// Read kubeconfig file on this runtime by default look sunder ~/.kube for config file
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	clusterurl := config.Host
	// Ask for confirmtation to proceed in creating kubeconfig file
	var proceed bool = false

	prompt := &survey.Confirm{
		Message: "You are going to create a kubeconfig file for cluster: " + clusterurl + "\n", Default: true,
	}
	survey.AskOne(prompt, &proceed, nil)
	if proceed {
		var nsusrlist []string
		api := clientset.CoreV1()
		nsinterface := api.Namespaces()
		nslist, err := nsinterface.List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, nsitem := range nslist.Items {
			//fmt.Println(nsitem.ObjectMeta.Name)
			nsusrlist = append(nsusrlist, nsitem.ObjectMeta.Name)
		}

		selectedns := ""
		prompt := &survey.Select{
			Message: "Select one of the follwing namespaces of this cluster",
			Options: nsusrlist,
		}
		survey.AskOne(prompt, &selectedns, nil)
		namespace = selectedns
		fmt.Println("selected naemspace is ", namespace)
		// Get secret name
		fmt.Printf("Getting credentials needed for kubeconfig \n")
		var sausrlist []string
		sainterface := api.ServiceAccounts(namespace)
		salist, err := sainterface.List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, saitem := range salist.Items {
			sausrlist = append(sausrlist, saitem.ObjectMeta.Name)
		}
		saname := ""
		prompt = &survey.Select{
			Message: "Select one of the follwing service accounts of this cluster",
			Options: sausrlist,
		}
		survey.AskOne(prompt, &saname, nil)
		sa, err := api.ServiceAccounts(namespace).Get(saname, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}
		secret := string(sa.Secrets[0].Name)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("got secret name")

		// Get ca data from secret
		fmt.Println("Getting certificate authority data")
		ca, err := api.Secrets(namespace).Get(secret, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}
		scrt := string(ca.Data["ca.crt"])
		sEnc := b64.StdEncoding.EncodeToString([]byte(scrt))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("got certificate authority data")

		// Get token value from secret
		fmt.Printf("Getting token \n")
		token := string(ca.Data["token"])
		fmt.Println("got token")

		// calling function that generates kubeconfig yaml file
		genconfig(clustername, sEnc, clusterurl, namespace, saname, token)
	}

}
