using ConsoleRPG.Weapons;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.InteropServices.Marshalling;
using System.Text;
using System.Threading.Tasks;
using static System.Net.Mime.MediaTypeNames;

namespace ConsoleRPG
{
    internal abstract class CombatCharacter : Character, IAttackModule, IDefendModule
    {
        // TODO : find a way to make this not protected and private, already caused one bug


        // when creating enemy types, use the setters, don't set default values. This is so the level can alter the values
        public IWeapon Weapon { get; set; }
        //public virtual int MaxHP { get { return (int)(maxHP * StatMultiplier); } set { maxHP = value; HP = maxHP; } }
        //public int HP { get; set; }
        //public int Armour { get { return (int)(armour * StatMultiplier); } set { armour = value; } }
        //public virtual int DodgeChance { get { return dodgeChance; } set { dodgeChance = value; } }
        public bool isDead => HP <= 0;
        // -0.1 to make it so level 1 has no bonuses
        public float StatMultiplier => (float)(1 + (level * 0.1 - 0.1));


        //// This property will reset the next attack bonus to 0 whenever
        //// its get  is used. This is because its assuming that once the bonus
        //// has been grabbed, it has been applied
        //public int NextAttackBonus {  
        //    get { 
        //        int nReturn = nextAttackBonus;
        //        nextAttackBonus = 0;
        //        return nReturn;
        //    }
        //    set {
        //        nextAttackBonus = value;
        //    }
        //}


        //public CombatCharacter(string characterName, int level) : base(characterName)
        //{
        //    this.level = level;
        //    Weapon = new Fists();
        //}



        // This method will calculate the damage that will be attacked with
        // It is given a weapon so future functionality can change which weapon
        // Is used to attack
        protected int CalculateDamage(IWeapon weapon, bool isCritical)
        {
            int damage = (int)((weapon.BaseDamage + NextAttackBonus) * StatMultiplier);

            if (isCritical)
            {
                damage = (int)(weapon.CriticalBonus * damage);
            }

            return damage;
        }

        protected virtual bool IsCritical(IWeapon weapon)
        {
            float critChance = weapon.CriticalChance;

            bool isCritical = false;

            Random rand = new Random();
            if (rand.NextDouble() < critChance)
            {
                isCritical = true;
            }

            return isCritical;

        }

        protected virtual int GetDamage(IWeapon weapon)
        {
            return 
        }


        public void TakeDamage(int nDamage, bool bBlockable)
        {
            HP -= bBlockable ? nDamage - Armour : nDamage;
        }
    }
}
