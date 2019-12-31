
def main():
    filename = "/Users/a123/prj/gopath/src/github.com/gkany/graphSDK/tests/myTest/op.txt"
    with open(filename, 'r') as file:
        for line in file:
            # print(line)
            line = line.strip()
            if len(line) > 0:
                tokens = line.split('_')
                new_line = "OperationType"
                for token in tokens:
                    new_line = new_line + token.title()
                print(new_line)

            # file.write(chains_data)
    file.close()


if __name__ == "__main__":
    main() 
