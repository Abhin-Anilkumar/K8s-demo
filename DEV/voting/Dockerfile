FROM maven:3-eclipse-temurin-19-alpine AS build
WORKDIR /app
COPY . .
RUN mvn package
FROM openjdk:19-jdk-alpine3.16
WORKDIR /app
COPY --from=build /app/target/voting-0.0.1-SNAPSHOT.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
