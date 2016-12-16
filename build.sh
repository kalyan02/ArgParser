java -jar /usr/local/lib/antlr-4.6-complete.jar ArgParse.g4 \
    -long-messages \
    -o output_java
java -jar /usr/local/lib/antlr-4.6-complete.jar ArgParse.g4 \
    -long-messages \
    -Dlanguage=Go \
    -o output_go
javac output_java/ArgParse*.java

