def label = "k8s-slave-golang-${UUID.randomUUID().toString()[0..7]}"


// loading neccessary docker images
podTemplate( cloud: 'openshift', label: label, namespace: "altemista-jenkins", slaveConnectTimeout: 30, containers: [
    containerTemplate(name: 'go', image: 'tetafro/golang-gcc:1.10-alpine', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'dind', image: 'docker:18.05-dind', privileged: true, ttyEnabled: true)
  ],
    volumes: [
                    emptyDirVolume(mountPath: '/var/lib/docker', memory: false),
                    hostPathVolume(hostPath: '/usr/share/zoneinfo/Europe/Bucharest', mountPath: '/etc/localtime'),
                    hostPathVolume(hostPath: '/etc/docker/certs.d', mountPath: '/etc/docker/certs.d'),
                    hostPathVolume(hostPath: '/var/run.docker.sock', mountPath: '/var/run.docker.sock'),
                    hostPathVolume(hostPath: '/etc/docker', mountPath: '/etc/docker')

]) {

    node(label) {
      try{
          
            stage('Checkout') {
                  checkout scm
            }
      
            stage('test') {
                container('go') {
                      sh 'go test ./...'
                }
            }
            
            stage('build'){
                container('go') {
                        sh 'go build -o hello-go'
                }
            } 
      
      
      currentBuild.result = 'SUCCESS'
       } catch (exc) {
                    currentBuild.result = 'FAILURE'
                  
                    // In case that the build is aborted
                    if (exc instanceof java.lang.InterruptedException) {
                        currentBuild.result = 'ABORTED'
                    }
       } finally {
           // Optional: send hipchat notification
           // pipelineUtilities.sendBuildNotifications(hipChatRoom)
        }
        }
    }
