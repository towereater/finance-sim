# Stage 1: Gradle Build
FROM gradle:8.14.3-jdk21 AS builder

WORKDIR /app
COPY . .

RUN gradle build

# Stage 2: Final image
FROM openjdk:21

# Exposed server port required at build time
ARG SERVER_PORT
ARG JAR_FILE=xchanger-0.0.1-SNAPSHOT.jar

WORKDIR /app
COPY --from=builder /app/build/libs/${JAR_FILE} /app

EXPOSE ${SERVER_PORT}

ENTRYPOINT ["java", "-Dserver.port=${SERVER.PORT}", "-jar", "xchanger-0.0.1-SNAPSHOT.jar"]