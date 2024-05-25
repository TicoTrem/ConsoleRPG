// See https://aka.ms/new-console-template for more information
using ConsoleRPG;
using ConsoleRPG.Enemies;
using ConsoleRPG.Weapons;



Console.WriteLine("Hello, adventurer. What is your name?");
string userName = Console.ReadLine();
MainCharacter user = new MainCharacter(userName);
Console.WriteLine($"Well then {user.CharacterName}, lets get you on your way.");
Console.WriteLine($"We are in dire need of common swordsman, and you are not special. Take this.");

user.Weapon = new TrainingSword();
ConsoleUtils.DisplaySystemMessage("You obtained a training sword!");

ConsoleUtils.DisplayWeaponStats(user.Weapon);

Console.WriteLine("A bozo approaches, press enter to attack");

Console.ReadLine();

CombatCharacter enemy1 = new Bozo(1);
CombatCharacter enemy2 = new Bozo(1);

BattleArena firstBattle = new BattleArena(user, [enemy1, enemy2]);

// if the user won
//if (firstBattle.StartBattle())
//{
//    Console.WriteLine("Congrats dawg, you win");
//}

ConsoleUtils.DisplayScenario("You encounter a chest", "Open it", "Ignore it", "Pee on it");







