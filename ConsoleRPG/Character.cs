using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleRPG
{
    internal class Character
    {
        private string characterName;
        private List<string> catchPhrases = ["Oh nyoooo", "frick u"];

        public string CharacterName { get; set; }
        public List<string> CatchPhrases { get; set; }

        public Character(string characterName)
        {
            CharacterName = characterName;
        }

        protected void Say(string message)
        {
            Console.WriteLine($"{this.CharacterName}: {message}");
        }

        


        

    }
}
