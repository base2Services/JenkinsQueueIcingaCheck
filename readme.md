Setup & install

go build *.go
go get -u github.com/base2services/golang-jenkins
mv JenkinsQueueIcingaCheck /etc/icinga/commands/....

Usage:

./JenkinsQueueIcingaCheck --jenkins-url=http://<jenkinsURLgoeshere>/ --jenkins-user=<jenkinsusernamegoeshere> --jenkins-apitoken=<jenkinsapitokengoeshere> --queue-check
