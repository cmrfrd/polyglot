FROM iflavoursbv/sbt-openjdk-8-alpine:0.13.15
ENV SCALA_LINK=http://downloads.typesafe.com/scala/2.11.8/scala-2.11.8.tgz
RUN mkdir -p /scala &&    wget -O /tmp/scala.tgz     $SCALA_LINK &&    tar -C /scala --strip-components=1 -zxvf /tmp/scala.tgz && rm /tmp/scala.tgz
ENV PATH /scala/bin:$PATH
ENV SBT_OPTS=""
WORKDIR /
ADD build.sbt /build.sbt
RUN sbt update dependencyClasspath
ADD solution.scala /solution.scala
RUN ["bash", "-c", "scalac -cp $(find /root/.ivy2/cache/ -type f -name '*jar' -printf '%p:') solution.scala"]
CMD ["bash", "-c", "scala solution.Main"]