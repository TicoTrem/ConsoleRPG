using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleRPG.Weapons
{
    internal class Fists : IWeapon
    {

        private string name = "Fists" /* + getRandomName */;
        private int baseDamage = 1;
        private IWeapon.DamageType damageType = IWeapon.DamageType.Regular;
        private string descriptorWord = "fragile";
        private string attackMessage = "That didn't do much of anything...";
        public float CriticalChance => 0.1f;
        public float CriticalBonus => 1.2f;

        public string Name { get { return name; } }
        public int BaseDamage { get { return baseDamage; } }
        public IWeapon.DamageType WeaponDamageType { get { return damageType; } }
        public string DescriptorWord { get { return descriptorWord; } }
        public string AttackMessage { get { return attackMessage; } }

    }
}
