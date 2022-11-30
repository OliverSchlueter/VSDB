package de.oliver;

import de.oliver.vsdb.Database;

import java.io.IOException;
import java.net.URISyntaxException;

public class Main {
    public static void main(String[] args) throws IOException, URISyntaxException {

        Database database = new Database("localhost", 80);
    }
}