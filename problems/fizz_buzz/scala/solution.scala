package main
object Main extends App {
  var string = "" : String

  (1 until 151)
  .map { i => (i%3, i%5, i%7) }
  .foreach { i =>
    i match {
      case (0,0,0) => string += "fizzbuzzbaz"
      case (0,_,0) => string += "fizzbaz"
      case (_,0,0) => string += "buzzbaz"
      case (0,0,_) => string += "fizzbuzz"
      case (0,_,_) => string += "fizz"
      case (_,0,_) => string += "buzz"
      case (_,_,0) => string += "baz"
      case _ => null
    }
  }
  println(string)
}
