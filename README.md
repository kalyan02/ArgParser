## ArgParse

Simple grammar/parser for parsing command line arguments.

Parser is generated using [Antlr v4](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)

## Usage

Assuming you have already installed Antlr, just run

    ~$ ./build.sh

this generates parsers for Go and java.

To try it out, run

    ~$ ./run.sh

## features

Grammar can identify

    --boolean
    --key value
    --key=value
    --key "a=b"
    --key="string value"

It doesn't work for

    --multi key1=val1 key2=val2

## Output

output visualized using antlr4 test rig.

![output](https://raw.githubusercontent.com/kalyan02/ArgParser/master/antlr4_parse_tree.png)

## License

BSD
