# Build stage
FROM maven:3.9-eclipse-temurin-21 AS builder

WORKDIR /build

COPY pom.xml .
COPY src ./src

RUN mvn clean package -DskipTests

# Runtime stage
FROM eclipse-temurin:21-jre

WORKDIR /app

COPY --from=builder /build/target/*.jar nm-api.jar

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "nm-api.jar"]