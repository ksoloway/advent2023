if __name__ == '__main__':
    file = open('input.txt')
    total_sum = 0
    for next_line in file:
        l, r = 0, len(str(next_line))-1
        
        while l < len(next_line):
            if next_line[l].isdigit():
                break
            l += 1
        
        while r >= 0:
            if next_line[r].isdigit():
                break
            r -= 1
        total_sum += 10*int(next_line[l]) + int(next_line[r])
    print(total_sum)
        