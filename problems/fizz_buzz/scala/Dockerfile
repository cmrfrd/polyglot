FROM hseeberger/scala-sbt:8u141-jdk_2.12.3_0.13.16
ENV SBT_OPTS=""
ADD solution.scala /solution.scala
WORKDIR /
CMD ["bash", "-c", "sbt --error 'set showSuccess := false' run"]
