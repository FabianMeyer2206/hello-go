pipeline {
    agent { docker { image 'golang:latest' } }
    stages {
        stage('build'){
            steps {
                sh '''
                    go build -o hello-go
                    ./hello-go
                '''
            }
        }
    }
}
