package protocol;

import java.io.IOException;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;
public class Packet {
    private final int id;
    private final int type;
    private final byte[] payload;

    public Packet(int id, int type, byte[] payload) {
        this.id = id;
        this.type = type;
        this.payload = payload;
    }

    public Packet(ByteBuffer buffer) throws IOException {
        int size = buffer.getInt();

        this.id = buffer.getInt();
        this.type = buffer.getInt();
        this.payload = new byte[size - 10];
        buffer.get(this.payload);

        buffer.get(new byte[2]);
    }

    public ByteBuffer toBuffer() {
        ByteBuffer buffer = ByteBuffer.allocate(this.payload.length + 14);
        buffer.order(ByteOrder.LITTLE_ENDIAN);

        buffer.putInt(this.payload.length + 10);
        buffer.putInt(this.id);
        buffer.putInt(this.type);
        buffer.put(this.payload);

        buffer.put((byte) 0);
        buffer.put((byte) 0);

        buffer.flip();
        return buffer;
    }

    public int getId() {
        return id;
    }

    public int getType() {
        return type;
    }

    public byte[] getPayload() {
        return payload;
    }
}
