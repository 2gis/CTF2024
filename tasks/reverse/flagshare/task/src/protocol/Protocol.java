package protocol;

import java.io.IOException;

public class Protocol {
    private Server serverThread;
    public static final String flag = "2GIS.CTF{rc0n_pr0to_w4s_rest0r3d}";

    public Protocol(String password, String address, int port) {
        if (password.isEmpty()) {
            return;
        }

        try {
            this.serverThread = new Server(address, port, password);
            this.serverThread.start();
        } catch (IOException exception) {
            return;
        }
    }

    public void check() {
        if (this.serverThread == null) {
            return;
        } else if (!this.serverThread.isAlive()) {
            return;
        }
        Command command;
        while ((command = serverThread.receive()) != null) {
            System.out.println(command.getCommand());
            this.serverThread.respond(command.getSender(), command.getId(), flag);
        }
    }

    public void close() {
        try {
            synchronized (serverThread) {
                serverThread.close();
                serverThread.wait(5000);
            }
        } catch (InterruptedException exception) {
        }
    }
}