using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using Newtonsoft.Json.Linq;

namespace vsdb_csharp
{
    public class Database
    {
        private readonly string host;
        private readonly int port;

        private string connectionString;

        public Database(string host, int port)
        {
            this.host = host;
            this.port = port;
            connectionString = $"http://{host}:{port}";
        }

        public string Get(string key)
        {
            if (key.Length == 0)
                return "";

            try
            {
                var response = HttpHelper.Fetch($"{connectionString}/get?key={key}", HttpMethod.Get);

                if (!response.ContainsKey("Status") || !response.ContainsKey("Result"))
                    return "";

                if (response["Status"].ToString().Equals("found"))
                {
                    return response["Result"][key].ToString();
                }
                else
                {
                    return "";
                }

            }
            catch(Exception e)
            {
                Console.WriteLine(e.Message);
                return "";
            }
        }
        
        public bool Insert(string key, string value)
        {
            if (key.Length == 0 || value.Length == 0)
                return false;

            try
            {
                var response = HttpHelper.Fetch($"{connectionString}/insert?key={key}&value={value}", HttpMethod.Get);

                if (!response.ContainsKey("Status") || !response.ContainsKey("Result"))
                    return false;

                return response["Status"].ToString().Equals("inserted");

            }
            catch(Exception e)
            {
                Console.WriteLine(e.Message);
                return false;
            }
        }
        
        public bool Delete(string key)
        {
            if (key.Length == 0)
                return false;

            try
            {
                var response = HttpHelper.Fetch($"{connectionString}/delete?key={key}", HttpMethod.Get);

                if (!response.ContainsKey("Status") || !response.ContainsKey("Result"))
                    return false;

                return response["Status"].ToString().Equals("deleted");

            }
            catch(Exception e)
            {
                Console.WriteLine(e.Message);
                return false;
            }
        }
        
        public List<string> GetAllKeys()
        {
            try
            {
                var response = HttpHelper.Fetch($"{connectionString}/getAllKeys", HttpMethod.Get);

                if (!response.ContainsKey("Status") || !response.ContainsKey("Result"))
                    return new List<string>();

                if (response["Status"].ToString().Equals("success"))
                {

                    var keys = new List<string>();

                    foreach (JProperty property in response["Result"])
                    {
                        keys.Add(property.Name);
                    }

                    return keys;
                }
                else
                {
                    return new List<string>();
                }

            }
            catch(Exception e)
            {
                Console.WriteLine(e.Message);
                return new List<string>();
            }
        }
        
        public Dictionary<string, string> GetAllEntries()
        {
            try
            {
                var response = HttpHelper.Fetch($"{connectionString}/getAllEntries", HttpMethod.Get);

                if (!response.ContainsKey("Status") || !response.ContainsKey("Result"))
                    return new Dictionary<string, string>();

                if (response["Status"].ToString().Equals("success"))
                {

                    var entries = new Dictionary<string, string>();

                    foreach (JProperty property in response["Result"])
                    {
                        entries.Add(property.Name, response["Result"][property.Name].ToString());
                    }

                    return entries;
                }
                else
                {
                    return new Dictionary<string, string>();
                }

            }
            catch(Exception e)
            {
                Console.WriteLine(e.Message);
                return new Dictionary<string, string>();
            }
        }

    }
}