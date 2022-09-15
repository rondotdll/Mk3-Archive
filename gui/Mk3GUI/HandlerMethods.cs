using System;
using MahApps.Metro.IconPacks;
using System.Collections.Generic;


namespace Mk3GUI
{
    internal class HandlerMethods
    {
        private static PackIconFontAwesomeKind[] ComputeStarValues(double starsValue)
        {
            PackIconFontAwesomeKind[] result = new PackIconFontAwesomeKind[5];

            for (int i = 0; i < 5; i++)
            {
                double difference = i + 1 - starsValue;
                if (difference > 0 && difference < 1)
                {
                    double remainder = starsValue % 1;

                    if (remainder >= 0.3 && remainder <= 0.7)
                        result[i] = PackIconFontAwesomeKind.StarHalfAltSolid;
                    else if (remainder > 0.7)
                        result[i] = PackIconFontAwesomeKind.StarSolid;

                } else if (i+1 <= starsValue)
                {
                    result[i] = PackIconFontAwesomeKind.StarSolid;
                } else
                {
                    result[i] = PackIconFontAwesomeKind.StarRegular;
                }

            }

            return result;
        }

        private static void SetStarOpacity(PackIconFontAwesome star)
        {
            switch (star.Kind)
            {
                case PackIconFontAwesomeKind.StarHalfAltSolid:
                    star.Opacity = 0.6;
                    break;
                case PackIconFontAwesomeKind.StarRegular:
                    star.Opacity = 0.2;
                    break;
                default:
                    star.Opacity = 1;
                    break;
            }
        }

        public static void UpdateDetectionRating(MainWindow Main)
        {

            double HighestValue = 0;

            foreach (KeyValuePair<string, FeatureProperties> Feature in Globals.FeatureSet)
            {
                if (Feature.Value.Enabled && (Feature.Value.Impact > HighestValue))
                {
                    HighestValue = Feature.Value.Impact;
                }
            }


            // stars are inversely proportional to the rating, high detection rating = low star count :(
            double starsValue = 5 - (HighestValue * 10);

            string color;

            if (HighestValue * 200 <= 33)
            {
                color = "green";
            } else if (33 < HighestValue*200 && HighestValue*200 <= 66)
            {
                color = "yellow";
            }
            else
            {
                color = "red";
            }
            Animations.FadeToColor(Main.DetectionPercentage, color, Math.Abs(Globals.DetectionRating - HighestValue) * 3);

            // In addition to modifying the star's "fullness" also modify the opacity.
            // Full-Star :   1.0
            // Half-Star :   0.6
            // Empty-Star :  0.2

            PackIconFontAwesomeKind[] stars = ComputeStarValues(starsValue);

            Main.FirstStar.Kind = stars[0];
            SetStarOpacity(Main.FirstStar);
            Main.SecondStar.Kind = stars[1];
            SetStarOpacity(Main.SecondStar);
            Main.ThirdStar.Kind = stars[2];
            SetStarOpacity(Main.ThirdStar);
            Main.FourthStar.Kind = stars[3];
            SetStarOpacity(Main.FourthStar);
            Main.FifthStar.Kind = stars[4];
            SetStarOpacity(Main.FifthStar);

            Globals.DetectionRating = HighestValue;
        }
    }
}
