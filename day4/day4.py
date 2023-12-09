from collections import defaultdict
if __name__ == '__main__':
    file = open('input.txt')
    total_sum = 0
    card_dict = defaultdict(lambda: 1)
    for line in file:
        game, line = line.strip().split(':')
        game = int(game.strip().split(' ')[-1])
        line = line.strip().split('|')
        winning, hand = line[0], line[1]
        winning = [ x for x in winning.strip().split(' ') if x]
        hand = [x for x in hand.strip().split(' ') if x]
        matching = 0
        for num in winning:
            if num in hand:
                matching += 1
        for i in range(1,matching+1):
            card_dict[i+game] += card_dict[game]
        total_sum += card_dict[game]
    print(total_sum)
            
