class RecentFile {
  final String? icon, title, date, size;

  RecentFile({this.icon, this.title, this.date, this.size});
}

List demoRecentFiles = [
  RecentFile(
    icon: "assets/icons/xd_file.svg",
    title: "Temp. Sensing",
    date: "02-09-2021",
    size: "1.3KB",
  ),
  RecentFile(
    icon: "assets/icons/Figma_file.svg",
    title: "Turn on sprinkler",
    date: "02-09-2021",
    size: "-",
  ),
  RecentFile(
    icon: "assets/icons/doc_file.svg",
    title: "Monitor Video",
    date: "02-09-2021",
    size: "32.5mb",
  ),
  RecentFile(
    icon: "assets/icons/google_drive.svg",
    title: "Syncronization",
    date: "02-09-2021",
    size: "250.7mb",
  ),
];
