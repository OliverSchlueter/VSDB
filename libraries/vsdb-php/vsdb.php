<?php

function vsdb_get($url, $key): string{
    try {
        $response = file_get_contents("$url/get?key=$key");

        if(!$response){
            return "";
        }

        $json = json_decode($response);

        if(isset($json->Status) && $json->Status == "found"){
            return $json->Result->$key;
        } else {
            return "";
        }

    } catch (Error | Exception){
        return "";
    }
}

function vsdb_insert($url, $key, $value): bool{
    try {
        $response = file_get_contents("$url/insert?key=$key&value=$value");

        if(!$response){
            return false;
        }

        $json = json_decode($response);

        if(isset($json->Status) && $json->Status == "inserted"){
            return true;
        } else {
            return false;
        }

    } catch (Error | Exception){
        return false;
    }
}

function vsdb_delete($url, $key): bool{
    try {
        $response = file_get_contents("$url/delete?key=$key");

        if(!$response){
            return false;
        }

        $json = json_decode($response);

        if(isset($json->Status) && $json->Status == "deleted"){
            return true;
        } else {
            return false;
        }

    } catch (Error | Exception){
        return false;
    }
}

function vsdb_get_all_keys($url): array|bool{
    try {
        $response = file_get_contents("$url/getAllKeys");

        if(!$response){
            return false;
        }

        $json = json_decode($response);

        if(isset($json->Status) && $json->Status == "success"){
            return array_keys(get_object_vars($json->Result));
        } else {
            return false;
        }

    } catch (Error | Exception){
        return false;
    }
}

function vsdb_get_all_entries($url): array|bool{
    try {
        $response = file_get_contents("$url/getAllEntries");

        if(!$response){
            return false;
        }

        $json = json_decode($response);

        if(isset($json->Status) && $json->Status == "success"){
            return (array) $json->Result;
        } else {
            return false;
        }

    } catch (Error | Exception){
        return false;
    }
}