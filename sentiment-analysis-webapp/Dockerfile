# Step : Test and package
FROM maven:3.5.3-jdk-8-alpine as build
WORKDIR /build
COPY pom.xml .
RUN mvn dependency:go-offline -B

COPY src/ /build/src/
RUN mvn package

FROM openjdk:8-jre-alpine
COPY --from=build /build/target/sentiment-analysis-web-0.0.1-SNAPSHOT.jar /app/sentiment-analysis-web-0.0.1-SNAPSHOT.jar
# Environment Variable that defines the endpoint of sentiment-analysis python api.
ENV SA_LOGIC_API_URL http://localhost:5000

EXPOSE 8080
CMD ["java", "-jar", "app/sentiment-analysis-web-0.0.1-SNAPSHOT.jar", "--sa.logic.api.url=${SA_LOGIC_API_URL}"]