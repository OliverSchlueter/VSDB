package de.oliver.vsdb;

import java.io.IOException;
import java.net.URISyntaxException;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;

public class Database {

    private final String host;
    private final int port;

    private final String connectionString;

    public Database(String host, int port) {
        this.host = host;
        this.port = port;
        this.connectionString = "http://" + host + ":" + port;
    }

    /**
     * @return the value
     */
    public String get(String key){
        if(key.length() == 0){
            return "";
        }

        try {
            String rawResponse = HttpHelper.httpRequest(connectionString + "/get?key=" + key);
            VSDBResponse response = HttpHelper.gson.fromJson(rawResponse, VSDBResponse.class);
            if(response.Status.equals("found")){
                return response.Result.get(key);
            } else {
                return "";
            }
        } catch (IOException | URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    public int getAsInt(String key){
        return Integer.parseInt(get(key));
    }

    public long getAsLong(String key){
        return Long.parseLong(get(key));
    }

    public float getAsFloat(String key){
        return Float.parseFloat(get(key));
    }

    public double getAsDouble(String key){
        return Double.parseDouble(get(key));
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

    public boolean insert(String key, int value){
        return insert(key, String.valueOf(value));
    }

    public boolean insert(String key, long value){
        return insert(key, String.valueOf(value));
    }

    public boolean insert(String key, float value){
        return insert(key, String.valueOf(value));
    }

    public boolean insert(String key, double value){
        return insert(key, String.valueOf(value));
    }

    /**
     * @return true if deleted, false if not
     */
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

    public Set<String> getAllKeys(){
        try {
            String rawResponse = HttpHelper.httpRequest(connectionString + "/getAllKeys");
            VSDBResponse response = HttpHelper.gson.fromJson(rawResponse, VSDBResponse.class);
            return response.Result.keySet();

        } catch (IOException | URISyntaxException e) {
            throw new RuntimeException(e);
        }
    }

    public Map<String, String> getAllEntries(){
        try {
            String rawResponse = HttpHelper.httpRequest(connectionString + "/getAllEntries");
            VSDBResponse response = HttpHelper.gson.fromJson(rawResponse, VSDBResponse.class);

            return response.Result;

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
