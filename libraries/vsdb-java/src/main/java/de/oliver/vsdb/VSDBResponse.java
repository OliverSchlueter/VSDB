package de.oliver.vsdb;

import java.util.Map;

public class VSDBResponse {

    public String Status;
    public Map<String, String> Result;

    public VSDBResponse(String Status, Map<String, String> Result) {
        this.Status = Status;
        this.Result = Result;
    }
}
