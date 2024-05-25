using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleRPG
{
    internal class BattleArena
    {


        private MainCharacter mainCharacter;
        private List<CombatCharacter> enemyList;

        private const int baseXPToGrant = 100;
        private const int bonusXPPerLevel = 10;




        internal BattleArena(MainCharacter mainCharacter, List<CombatCharacter> enemyList)
        {
            this.mainCharacter = mainCharacter;
            this.enemyList = enemyList;
        }

        // This method will take in two CombatCharacters, and make them battle
        // At the end, true will be returned if main character won and false will be returned
        // if the enemies won
        internal bool StartBattle()
        {

            // the list is a value copy, the enemies inside are a reference copy
            List<CombatCharacter> enemyListCopy = new List<CombatCharacter>(enemyList);
            while (enemyList.Count > 0 && mainCharacter.HP >= 0)
            {
                Thread.Sleep(1000);
                MCTakeTurn();
                
                foreach (CombatCharacter enemy in enemyList)
                {
                    if (enemy.HP > 0 && mainCharacter.HP > 0)
                    {
                        Thread.Sleep(1000);
                        enemy.Attack(mainCharacter);
                    }
                }
            }

            bool mcWon = mainCharacter.HP > 0;
            if (mcWon)
            {
                GrantExperience(enemyListCopy);
            }

            return mcWon;
        }


        internal void MCTakeTurn()
        {
            DisplayEnemies();
            CombatCharacter enemy = null;
            bool validInput = false;
            do
            {
                int enemySelection;
                Console.WriteLine("Select an enemy to attack (Enter their number)");
                string input = Console.ReadLine();

                if (int.TryParse(input, out enemySelection)) // if successful, puts the parsed value in to enemySelection
                {
                    enemy = enemyList[enemySelection - 1];
                    validInput = enemySelection >= 1 && enemySelection <= enemyList.Count;
                }

                
            } while (!validInput); // will put the parsed

            // chose to keep the console.writeline in the CombatCharacter attack method
            // because this BattleArena class handles the battle, turns, info displaying, and choices while
            // the CombatCharacter can perform the action of attacking. Things like
            // Critical hit are also in that combat character class for good reason
            mainCharacter.Attack(enemy!); // -1 to convert to 0 based index

            // coudl put this in CombatCharacter or Maincharacter,
            // but i think its best that The character just handles that it can
            // attack, and BattleArena controls flow of the enemies inside
            // besides, we couldnt remove them from the list then

            // checks if enemy that was attacked is dead
            if (enemy!.isDead)
            {
                Random rand = new Random();
                //string catchPhrase = enemy.CatchPhrases[rand.Next(0, enemy.CatchPhrases.Count)];
                Console.WriteLine($"It was a killing blow. ");
            }
        }


        // called when the main character wins the battle
        // grants experience for each enemy killed based on enemy level
        internal void GrantExperience(List<CombatCharacter> originalEnemyList)
        {
            foreach (CombatCharacter enemy in originalEnemyList)
            {
                mainCharacter.increaseXP((enemy.Level * baseXPToGrant) + baseXPToGrant);
            }
        }





        internal void DisplayEnemies()
        {
            int enemyNumber = 1;
            foreach (CombatCharacter enemy in enemyList)
            {
                Console.WriteLine($"Enemy {enemyNumber++}: {enemy.CharacterName} HP {enemy.HP}/{enemy.MaxHP} Weapon {enemy.Weapon.Name}");
            }
        }

        


    }
}
