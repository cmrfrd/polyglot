FROM hseeberger/scala-sbt:8u141-jdk_2.12.3_0.13.16
USER root
ENV SBT_VERSION 0.13.15
ENV DEST_DIR /usr/local
ENV SBT_HOME ${DEST_DIR}/sbt
ENV PATH ${PATH}:${SBT_HOME}/bin
ENV SBT_OPTS="-Dsbt.override.build.repos=true"
WORKDIR /root
RUN mkdir -p project &&     echo "scalaVersion := \"2.11.8\"" > build.sbt &&     echo "sbt.version=${SBT_VERSION}" > project/build.properties &&     echo "case object Temp" > Temp.scala &&     sbt compile &&     rm -r project && rm build.sbt && rm Temp.scala && rm -r target
ENV SCALA_LINK=http://downloads.typesafe.com/scala/2.11.8/scala-2.11.8.tgz
RUN mkdir -p /scala &&    wget -O /tmp/scala.tgz     $SCALA_LINK &&    tar -C /scala -zxvf /tmp/scala.tgz && rm /tmp/scala.tgz
ENV PATH=/scala/scala-2.11.8/bin:$PATH
ADD build.sbt /root/build.sbt
RUN sbt update dependencyClasspath
ADD solution.scala /root/solution.scala
RUN ["bash", "-c", "scalac -cp $(find /root/.ivy2/cache/ -type f -name '*jar' -printf '%p:') solution.scala"]
CMD ["scala", "solution.Main", "|", "grep -v warn"]
