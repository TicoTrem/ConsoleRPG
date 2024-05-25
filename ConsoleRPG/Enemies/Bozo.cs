using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleRPG.Enemies
{
    internal class Bozo : CombatCharacter
    {
        
        public Bozo(int level) : base("Bozo", level)
        {
            MaxHP = 20;
            dodgeChance = 0;
            Armour = 0;
        }
    }
}
