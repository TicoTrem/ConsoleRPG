namespace ConsoleRPG
{
    internal interface IDefendModule
    {
        public int HP { get; set; }
        public int MaxHP { get; set; }
        public float DodgeChance { get; set; }
        public int Armour { get; set; }
        public void TakeDamage(int damageToTake, bool isBlockable);
    }
}