pipeline {
    agent { docker { image 'golang:latest' } }
    stages {
        stage('test') {
            steps {
                sh '''
                    go test ./...
                '''
            }
        }
        stage('build'){
            steps {
                sh '''
                    go build -o hello-go
                '''
            }
        }
    }
}
