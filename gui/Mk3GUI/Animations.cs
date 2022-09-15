using System;
using System.Windows.Controls;
using System.Windows.Media;
using System.Windows.Media.Animation;

namespace Mk3GUI
{
    public static class Animations
    {
        public static DoubleAnimation FadeOutPartial = new DoubleAnimation()
        {
            EasingFunction = Globals.SmoothEase,
            From = 1,
            To = 0.45,
            Duration = TimeSpan.FromSeconds(.25),
            FillBehavior = FillBehavior.HoldEnd,
        };

        public static DoubleAnimation FadeInPartial = new DoubleAnimation()
        {
            EasingFunction = Globals.SmoothEase,
            From = 0.45,
            To = 1,
            Duration = TimeSpan.FromSeconds(.25),
            FillBehavior = FillBehavior.HoldEnd,
        };

        public static DoubleAnimation CheckBoxClickDown = new DoubleAnimation()
        {
            EasingFunction = Globals.SmoothEase,
            From = 0,
            To = -25,
            Duration = TimeSpan.FromSeconds(.15),
            FillBehavior = FillBehavior.HoldEnd,
        };

        public static DoubleAnimation CheckBoxClickDown2 = new DoubleAnimation()
        {
            EasingFunction = Globals.SmoothEase,
            From = 0,
            To = 25,
            Duration = TimeSpan.FromSeconds(.15),
            FillBehavior = FillBehavior.HoldEnd,
        };

        public static ColorAnimation FadeColorGreenIn = new ColorAnimation()
        {
            EasingFunction = Globals.SmoothEase,
            From = Color.FromArgb(13, 255, 255, 255),
            To = Color.FromArgb(204, 1, 230, 179),
            Duration = TimeSpan.FromSeconds(.15),
            FillBehavior = FillBehavior.HoldEnd
        };

        public static ColorAnimation FadeColorYellowIn = new ColorAnimation()
        {
            EasingFunction = Globals.SmoothEase,
            From = Color.FromArgb(13, 255, 255, 255),
            To = Color.FromArgb(204, 255, 225, 77),
            Duration = TimeSpan.FromSeconds(.15),
            FillBehavior = FillBehavior.HoldEnd
        };

        public static ColorAnimation FadeColorRedIn = new ColorAnimation()
        {
            EasingFunction = Globals.SmoothEase,
            From = Color.FromArgb(13, 255, 255, 255),
            To = Color.FromArgb(204, 255, 51, 119),
            Duration = TimeSpan.FromSeconds(.15),
            FillBehavior = FillBehavior.HoldEnd
        };

        public static void FadeToColor(Label target, string color, double duration)
        {
            ColorAnimation FadeToColor = new ColorAnimation()
            {
                EasingFunction = Globals.SmoothEase,
                From = (Color)ColorConverter.ConvertFromString(target.Foreground.ToString()),
                Duration = TimeSpan.FromSeconds(duration),
                FillBehavior = FillBehavior.HoldEnd
            };

            switch (color)
            {
                case "red":
                    FadeToColor.To = Color.FromArgb(204, 255, 51, 119);
                    break;

                case "yellow":
                    FadeToColor.To = Color.FromArgb(204, 255, 225, 77);
                    break;

                case "green":
                    FadeToColor.To = Color.FromArgb(204, 1, 230, 179);
                    break;
            }

            target.Foreground.BeginAnimation(SolidColorBrush.ColorProperty, FadeToColor);
        }

        public static void FadeColorOut(ref TextBlock target)
        {
            ColorAnimation FadeColorOut = new ColorAnimation()
            {
                EasingFunction = Globals.SmoothEase,
                From = (Color)ColorConverter.ConvertFromString(target.Foreground.ToString()),
                To = Color.FromArgb(13, 255, 255, 255),
                Duration = TimeSpan.FromSeconds(.15),
                FillBehavior = FillBehavior.HoldEnd
            };

            target.Foreground.BeginAnimation(SolidColorBrush.ColorProperty, FadeColorOut);
        }

        public static void DoCheckBoxBounce(Button target, bool inverted)
        {
            if (inverted)
            {
                DoubleAnimation CheckBoxClickUp2A = new DoubleAnimation()
                {
                    EasingFunction = Globals.SmoothEase,
                    From = 25,
                    To = -25,
                    Duration = TimeSpan.FromSeconds(.1),
                    FillBehavior = FillBehavior.HoldEnd,
                };

                DoubleAnimation CheckBoxClickUp2B = new DoubleAnimation()
                {
                    EasingFunction = Globals.SmoothEase,
                    From = -25,
                    To = 0,
                    Duration = TimeSpan.FromSeconds(.1),
                    FillBehavior = FillBehavior.HoldEnd,
                };

                CheckBoxClickUp2A.Completed += new EventHandler((object s, EventArgs args) => {
                    target.Dispatcher.Invoke(() => {
                        target.RenderTransform.BeginAnimation(RotateTransform.AngleProperty, CheckBoxClickUp2B);
                    });
                });

                target.RenderTransform.BeginAnimation(RotateTransform.AngleProperty, CheckBoxClickUp2A);
                return;
            }

            DoubleAnimation CheckBoxClickUpA = new DoubleAnimation()
            {
                EasingFunction = Globals.SmoothEase,
                From = -25,
                To = 25,
                Duration = TimeSpan.FromSeconds(.1),
                FillBehavior = FillBehavior.HoldEnd,
            };

            DoubleAnimation CheckBoxClickUpB = new DoubleAnimation()
            {
                EasingFunction = Globals.SmoothEase,
                From = 25,
                To = 0,
                Duration = TimeSpan.FromSeconds(.1),
                FillBehavior = FillBehavior.HoldEnd,
            };
            
            CheckBoxClickUpA.Completed += new EventHandler((object s, EventArgs args) => {
                target.Dispatcher.Invoke(() => {
                    target.RenderTransform.BeginAnimation(RotateTransform.AngleProperty, CheckBoxClickUpB);
                });
            });

            target.RenderTransform.BeginAnimation(RotateTransform.AngleProperty, CheckBoxClickUpA);
        }
    }
}
