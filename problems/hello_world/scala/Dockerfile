FROM iflavoursbv/sbt-openjdk-8-alpine:0.13.15
ENV SBT_OPTS=""
ADD solution.scala /solution.scala
WORKDIR /
CMD ["bash", "-c", "sbt --error 'set showSuccess := false' run"]
