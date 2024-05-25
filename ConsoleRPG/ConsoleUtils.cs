using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleRPG
{
    internal class ConsoleUtils
    {
        internal static void DisplayWeaponStats(IWeapon weapon)
        {
            string weaponInfo = $"Weapon: {weapon.Name}\nDamage: {weapon.BaseDamage}\nDamage Type: {weapon.WeaponDamageType}";
            DisplaySystemMessage(weaponInfo);
        }

        internal static void DisplaySystemMessage(string message)
        {
            Console.WriteLine("--------------------------------------------------------------------------------------");
            Console.WriteLine($"* {message} *");
            Console.WriteLine("--------------------------------------------------------------------------------------");
        }

        internal static void DisplayScenario(string scenario, params string[] options)
        {
            Console.WriteLine(scenario + "\n");
            for (int i = 0; i < options.Length; i++) 
            {
                Console.WriteLine($"{i + 1}) {options[i]}");
            }
            Console.WriteLine("\n\no) Options\ts) Save Game\tl) Load Game\tq) Quit");
        }



        //internal static int GrantExperience(MainCharacter mc)
        //{

        //}
    }

    

}
