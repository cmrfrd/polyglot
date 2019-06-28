logLevel := Level.Error
scalaVersion := "2.11.8"
libraryDependencies ++= Seq("org.scalanlp" %% "breeze" % "0.13.2")
mainClass in (Compile,run) := Some("main.Main")
