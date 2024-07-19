import protocol.Protocol;

public class Main {

    public static void main(String[] args) throws InterruptedException {
        Protocol protocol = new Protocol("thisisrealpasswordyoucanuseitforyourexploit", "127.0.0.1", 7125);
        while(true) {
            protocol.check();
            Thread.sleep(250);
        }
    }
}