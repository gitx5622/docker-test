pipeline {
        environment {
            registry = "gits5622/docker-test-backend"
            registryCredential = 'docker-hub'
            dockerImage = ''
            // dockermigrate = docker.build("app -f Dockerfile-migrate")
            }
        agent any
        stages {
                stage('Cloning our Git') {
                    steps {
                    git 'https://github.com/gitx5622/docker-test.git'
                    }
                }
                stage('Building our image') {
                    steps{
                        script {
                        dockerImage = docker.build registry + ":$BUILD_NUMBER"
                        }
                    }
                    post{
                        always{
                            echo "Running docker migrate with golang-migrate"
                        }
                        success{
                            // script {
                             //       dockermigrate
                           //  }
                           echo "Build Success"
                        }
                        failure{
                            echo "Failed"
                        }
                    }
                }
                stage('Deploy our image') {
                    steps{
                        script {
                        docker.withRegistry( '', registryCredential ) {
                        dockerImage.push()
                            }
                        }
                    }
        }

    }
}