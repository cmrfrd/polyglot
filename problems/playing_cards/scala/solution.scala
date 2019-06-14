package solution

import breeze.linalg._

object CardConstants {
  val suits = List("D","C","H","S")
  val values = (1 until 11).map(i => i.toString()) ++ List("J", "Q", "K")
}

// def cartesianProduct[T](inputs: List[List[T]]) List[List[T]] = inputs match {
//   case Nil => List(List())
//   case h :: t => h.flatMap(i -> cartesianproduct().map(i :: _) 
// }

object Main extends App {
  println("potato!")
}
