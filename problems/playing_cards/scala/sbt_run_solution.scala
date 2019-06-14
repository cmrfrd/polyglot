#!/usr/bin/env sbt --error 'set showSuccess := false'
/***
name := """problem_solution"""
mainClass := Some("main.Main")
libraryDependencies ++= Seq(
  "com.chuusai" %% "shapeless" % "2.3.3"
)
*/

import shapeless._

object CardConstants {
  val suits = List("D","C","H","S")
  val values = (1 until 11).map(i => i.toString()) ++ List("J", "Q", "K")
}

// def cartesianProduct[T](inputs: List[List[T]]) List[List[T]] = input match {
//   case Nil => List(Nil)
//   case h :: t => for () yield 
// }

val cc = CardConstants

println(cc.suits)
println(cc.values)
