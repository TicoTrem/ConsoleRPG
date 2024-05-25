namespace ConsoleRPG
{
    internal interface IAttackModule
    {
        public int NextAttackBonus { get; set; }
        public void Attack(IDefendModule target, IWeapon weapon, float statMultiplier);

    }
}