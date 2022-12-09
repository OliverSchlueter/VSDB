using System;

namespace vsdb_csharp
{
    internal class Program
    {
        public static void Main(string[] args)
        {
            var db = new Database("localhost", 80);
            foreach (var entry in db.GetAllEntries())
            {
                Console.WriteLine($"{entry.Key} : {entry.Value}");
            }
        }
    }
}