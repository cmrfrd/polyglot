from random import shuffle, choice, randint
from itertools import product
from collections import namedtuple

# create suits


class Deck:
    """
    Generic deck class
    """

    values = list(map(str, range(1, 11))) + ['J', 'Q', 'K']
    suits = ['D', 'C', 'H', 'S']
    Card = namedtuple("card", ['value', 'suit'])

    def __init__(self):
        self.cards = [self.Card(v, s)
                      for v, s in product(self.values, self.suits)]

    def __len__(self):
        return len(self.cards)

    def __iter__(self):
        return iter(self.cards)

    def __getitem__(self, i):
        return self.cards[i]

    def __setitem__(self, key, item):
        self.cards[key] = item


'''
############################################################
####################     solutions    ######################
############################################################
'''

# 1. Create deck of cards and print top 5
for i, card in enumerate(list(filter(lambda c: c.suit == "S", Deck().cards))):
    if i < 5:
        print(card.value+card.suit, end=" ")
print()

# 2. Shuffle cards and print 5 random cards
deck = Deck()
shuffle(deck)
for i, card in enumerate(deck):
    if i < 5:
        print (card.value+card.suit, end=" ")
print()

# 3. Sort a shuffled deck of cards


def get_numeric_value(card):
    return (Deck.values.index(card.value)) + \
           (20*Deck.suits.index(card.suit))


for card in sorted(deck, key=get_numeric_value):
    print (card.value+card.suit, end=" ")
print()

# 4. Compare random cards against eachother by [spades, hearts, clubs, diamonds] in that order
deck = Deck()
card1, card2 = choice(deck), choice(deck)
sign = "<" if get_numeric_value(card1) < get_numeric_value(card2) else ">"
print("%s %s %s" % (card1.value+card1.suit, sign, card2.value+card2.suit))

# 5. Cut a deck of cards
deck = Deck()
split = randint(0, len(deck))
deck[split:] + deck[:split]

# 6. simulate a simple game of "war"
new_deck = Deck()
shuffle(new_deck)
player1, player2 = new_deck[:len(new_deck)//2], new_deck[len(new_deck)//2:]

iters = 0
while (len(player1) != 0) and \
      (len(player2) != 0):

    iters += 1

    p1_card, p2_card = player1.pop(), player2.pop()

    if get_numeric_value(p1_card) == get_numeric_value(p2_card):
        player1.append(p1_card)
        player2.append(p2_card)
    elif get_numeric_value(p1_card) < get_numeric_value(p2_card):
        player2.append(p1_card)
        player2.append(p2_card)
    else:
        player1.append(p1_card)
        player1.append(p2_card)

print ("Player 1:", len(player1), "Player 2:", len(player2), "Iters: ", iters)
