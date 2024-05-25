namespace ConsoleRPG
{
    internal interface IWeapon
    {
        public string Name { get; }
        public int BaseDamage { get; }
        public DamageType WeaponDamageType { get; }
        public string DescriptorWord { get; }
        public string AttackMessage { get; }
        public float CriticalChance { get; }
        public float CriticalBonus { get; }
     
        

        internal enum DamageType
        {
            Regular,
            Poison,
            Bleeding,
            Electric
        }
    }
}