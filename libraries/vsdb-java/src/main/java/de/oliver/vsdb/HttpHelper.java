package de.oliver.vsdb;

import com.google.gson.Gson;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.URISyntaxException;
import java.net.URL;
import java.net.URLConnection;
import java.util.stream.Collectors;

public class HttpHelper {

    public static final Gson gson = new Gson();

    public static String httpRequest(String targetURL) throws IOException, URISyntaxException {
        URL url = new URL(targetURL);
        URLConnection connection = url.openConnection();
        connection.setUseCaches(false);
        connection.setDoOutput(true);
        BufferedReader reader = new BufferedReader(new InputStreamReader(connection.getInputStream()));
        return reader.lines().collect(Collectors.joining("\n"));
    }

}
