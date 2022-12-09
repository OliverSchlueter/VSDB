using System;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace vsdb_csharp
{
    public class HttpHelper
    {
        public static JObject Fetch(string url, HttpMethod method, string auth = "")
        {
            try
            {
                var client = new HttpClient();
                var request = new HttpRequestMessage(method, url);
                if (auth.Length > 0)
                {
                    request.Headers.Authorization =
                        new AuthenticationHeaderValue("Basic", Convert.ToBase64String(Encoding.UTF8.GetBytes(auth)));
                }

                var task = client.SendAsync(request);

                var result = task.Result.Content.ReadAsStringAsync().Result;

                var json = (JObject)JsonConvert.DeserializeObject(result);

                return json;
            }
            catch (Exception e)
            {
                Console.WriteLine(e.Message);
                return null;
            }
        }
    }
}
