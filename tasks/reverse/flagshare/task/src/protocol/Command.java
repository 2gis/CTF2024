package protocol;

import java.nio.channels.SocketChannel;
public class Command {

    private final SocketChannel sender;
    private final int id;
    private final String command;

    public Command(SocketChannel sender, int id, String command) {
        this.sender = sender;
        this.id = id;
        this.command = command;
    }

    public SocketChannel getSender() {
        return this.sender;
    }

    public int getId() {
        return this.id;
    }

    public String getCommand() {
        return this.command;
    }
}
