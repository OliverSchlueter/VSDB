package de.oliver.vsdb;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.HashMap;
import java.util.Map;

public class Database {

    private final String host;
    private final int port;

    private final String connectionString;

    public Database(String host, int port) {
        this.host = host;
        this.port = port;
        this.connectionString = "http://" + host + ":" + port;
    }

    public String get(String key){
        if(key.length() == 0){
            return "";
        }

        try {
            String rawResponse = HttpHelper.httpRequest(connectionString + "/get?key=" + key);
            VSDBResponse response = HttpHelper.gson.fromJson(rawResponse, VSDBResponse.class);
            if(response.Status.equals("found")){
                return response.Result;
            } else {
                return "";
            }
        } catch (IOException | URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    public boolean insert(String key, String value){
        if(key.length() == 0 || value.length() == 0){
            return false;
        }

        try {
            String rawResponse = HttpHelper.httpRequest(connectionString + "/insert?key=" + key + "&value=" + value);
            VSDBResponse response = HttpHelper.gson.fromJson(rawResponse, VSDBResponse.class);
            return response.Status.equals("inserted");
        } catch (IOException | URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    public boolean delete(String key){
        if(key.length() == 0){
            return false;
        }

        try {
            String rawResponse = HttpHelper.httpRequest(connectionString + "/delete?key=" + key);
            VSDBResponse response = HttpHelper.gson.fromJson(rawResponse, VSDBResponse.class);
            return response.Status.equals("deleted");
        } catch (IOException | URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    public String[] getAllKeys(){
        try {
            String rawResponse = HttpHelper.httpRequest(connectionString + "/getAllKeys");
            VSDBResponse response = HttpHelper.gson.fromJson(rawResponse, VSDBResponse.class);
            return response.Result.split(";");

        } catch (IOException | URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    public Map<String, String> getAllEntries(){
        try {
            String rawResponse = HttpHelper.httpRequest(connectionString + "/getAllEntries");
            VSDBResponse response = HttpHelper.gson.fromJson(rawResponse, VSDBResponse.class);
            Map<String, String> entries = new HashMap<>();
            String[] rawEntries = response.Result.split(";");
            for (String rawEntry : rawEntries) {
                String[] entry = rawEntry.split(":");
                entries.put(entry[0], entry[1]);
            }

            return entries;

        } catch (IOException | URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }


    public String getHost() {
        return host;
    }

    public int getPort() {
        return port;
    }
}
