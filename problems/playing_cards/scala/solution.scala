package solution

import scala.util.Random
import breeze.linalg._

object DeckData {
  val suits = List("D","C","H","S")
  val values = (1 until 11).map(i => i.toString()).toList ++ List("J", "Q", "K")
  val ordered_cards = Operations.cartesianProduct(List(suits, values)).view.zipWithIndex.map( c => (Map("suit"->c._1(0), "value"-> c._1(1)), c._2) ) toMap
}

object Operations {
  def cartesianProduct[T](inputs: List[List[T]]): List[List[T]] = inputs match {
    case Nil => List(List())
    case h :: t => h.flatMap(i => cartesianProduct(t).map(i :: _))
  }
  def printCard(card: Map[String, String]) = {
    printf(card.values.mkString("") ++ " ")
  }
}

class Deck {
  var cards = Operations.cartesianProduct(List(DeckData.suits, DeckData.values)).map(c => Map("suit"->c(0), "value"-> c(1)))


  def higher[T](card1: T, card2: T): Boolean = {
    DeckData.ordered_cards(card1.asInstanceOf[Map[String, String]] ) > DeckData.ordered_cards(card2.asInstanceOf[Map[String, String]])
  }

  def lower[T](card1: T, card2: T): Boolean = {
    DeckData.ordered_cards(card1.asInstanceOf[Map[String, String]]) < DeckData.ordered_cards(card2.asInstanceOf[Map[String, String]])
  }

  def shuffle() = {
    cards = Random.shuffle(cards)
    this
  }

  def sort() = {
    cards = quicksort(cards)
    this
  }

  def quicksort[T](x: List[T]): List[T] = {
    if (x.length <= 1) x
    else {
      val pivot = x(x.length / 2)
      quicksort(x filter (c => lower(c, pivot)) toList) ++
               (x filter (c => pivot == c) toList) ++
      quicksort(x filter (c => higher(c, pivot)) toList)
    }
  }
}


object Main extends App {

  // 1. Create a deck and print all sorted spades
  (new Deck().cards.filter(i => i("suit") == "S").take(5)) map Operations.printCard _
  println()

  //2. Shuffle cards and print 5 random cards
  val rand_deck = new Deck()
  rand_deck.shuffle().cards take 5 map Operations.printCard _
  println()

  //3. Sort a shuffled deck of cards
  //
  //
  rand_deck.sort().cards map Operations.printCard _
}
