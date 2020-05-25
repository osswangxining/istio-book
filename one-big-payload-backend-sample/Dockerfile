FROM openjdk:8-jdk-alpine
VOLUME /tmp
ADD target/one-big-payload-backend-sample-0.1.0.jar app.jar
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/app.jar"]