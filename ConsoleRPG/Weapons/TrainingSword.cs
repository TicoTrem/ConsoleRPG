using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection.Metadata.Ecma335;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleRPG.Weapons
{
    internal class TrainingSword : IWeapon
    {

        private string name = "Training Sword" /* + getRandomName */;
        private int baseDamage = 10;
        private IWeapon.DamageType damageType = IWeapon.DamageType.Regular;
        private string descriptorWord = "dull";
        private string attackMessage = "That didn't do much of anything...";


        public string Name { get { return name; } }
        public int BaseDamage { get { return baseDamage; } }
        public IWeapon.DamageType WeaponDamageType { get { return damageType; } }
        public string DescriptorWord { get { return descriptorWord; } }
        public string AttackMessage { get { return attackMessage; } }

        //public float CriticalChance => 0.05f;
        public float CriticalChance => 0.05f;
        public float CriticalBonus => 1.5f;


    }
}
