using System.Reflection.Emit;

namespace ConsoleRPG
{
    public abstract class ALeveler
    {
        public abstract int Level { get; set; }

        // will only be in the main character implementation
        //public void increaseXP(int amountToIncrease);
        //public void LevelUp();

        public float StatMultiplier => (float)(1 + (Level * 0.1 - 0.1));
    } 
}