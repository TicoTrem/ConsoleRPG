using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleRPG
{
    internal class MainCharacterAttackModule : IAttackModule
    {

        private ICharacterProperties characterProperties;
        public int NextAttackBonus { get; set; }

        private int statMultiplier;

        public MainCharacterAttackModule(ICharacterProperties characterProperties)
        {
            this.characterProperties = characterProperties;
        }

        public virtual void Attack(IDefendModule target, IWeapon weapon, double statMultiplier = 1)
        {
            
            bool isCritical = IsCritical(weapon);
            int damageToDeal = CalculateDamage(weapon, isCritical);
            target.TakeDamage(damageToDeal, true);
            Console.Write($"{characterProperties.CharacterName } used their {weapon.Name} and Dealt {damageToDeal} to {characterProperties.CharacterName}\n" +
                                $"{characterProperties} now has {target.HP}/{target.MaxHP} health ");
            if (isCritical)
            {
                Console.WriteLine("Critical Hit!");
            }
        }
















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

        protected int GetDamage(IWeapon weapon)
        {
            float damageMultiplier = (float)(damageLevel * 0.1);
            int baseDamage = (int)((weapon.BaseDamage + NextAttackBonus) * StatMultiplier);
            return (int)(baseDamage * damageMultiplier);
        }

        protected bool IsCritical(IWeapon weapon)
        {
            // increase by 2% per level 
            float critChance = weapon.CriticalChance + this.CriticalChance;

            bool isCritical = false;

            Random rand = new Random();
            Console.WriteLine(critChance);
            Console.WriteLine(rand.NextDouble());
            if (rand.NextDouble() < critChance)
            {
                isCritical = true;
            }

            return isCritical;

        }

        public void Attack(IDefendModule target, IWeapon weapon, float statMultiplier)
        {
            throw new NotImplementedException();
        }
    }
}
