using ConsoleRPG.Weapons;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection.Emit;
using System.Reflection.Metadata.Ecma335;
using System.Text;
using System.Threading.Tasks;


namespace ConsoleRPG
{
    internal class MainCharacter
    {
        // it can attack
        private IAttackModule attackModule;
        // it can defend
        private IDefendModule defendModule;
        // it can level up
        private ILevelUp levelUpModule;
            // just add the xp,  level, LevelUp, and increaseLevel methods?

        private int xp;

        private int maxHP;
        private int nextAttackBonus;
        private float dodgeChance;

        public int NextAttackBonus
        {
            get
            {
                int nReturn = nextAttackBonus;
                nextAttackBonus = 0;
                return nReturn;
            }
            set
            {
                nextAttackBonus = value;
            }
        }

        public int MaxHP { get { return (int)((maxHP + (healthLevel * 2)) * StatMultiplier); } set { maxHP = value; HP = MaxHP; } }
        public float DodgeChance { get { return dodgeChance + dodgeLevel; } set { dodgeChance = value; } }

        // just the critical chance added by the crit level, does not include weapon crit chance
        public float CriticalChance => (float)(critLevel * 0.01) * 2;


        private int healthLevel = 0;
        private int damageLevel = 0;
        private int dodgeLevel = 0;
        private int critLevel = 0;


        public void IncreaseHealthLevel() { healthLevel++;  }
        public void IncreaseDamageLevel() { damageLevel++;  }
        public void IncreaseDodgeLevel() { dodgeChance++;  }
        public void IncreaseCritLevel() { critLevel++;  }



        public void LevelUp() { level++; }
        public void increaseXP(int amountToIncrease)
        {
            xp += amountToIncrease;
            if (xp > 100)
            {
                LevelUp();
                xp = 0;
            }
        } 



        // null weapon will be assigned fists in CombatCharacter constructor
        public MainCharacter(string characterName) : base(characterName, 1)
        {
            MaxHP = 20;
            xp = 0;
        }



    }
}
