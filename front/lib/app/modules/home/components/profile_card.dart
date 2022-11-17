import 'package:flutter/material.dart';
import 'package:front/app/components/responsive.dart';
import 'package:front/colors.dart';

class ProfileCard extends StatelessWidget {
  const ProfileCard({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var dropdownBtnKey = GlobalKey();
    return Container(
        margin: const EdgeInsets.only(left: defaultPadding),
        padding: const EdgeInsets.symmetric(
          horizontal: defaultPadding,
          vertical: defaultPadding / 2,
        ),
        decoration: BoxDecoration(
          color: secondaryColor,
          borderRadius: const BorderRadius.all(Radius.circular(10)),
          border: Border.all(color: Colors.white10),
        ),
        child: SizedBox(
          height: 36,
          child: DropdownButton(
            items: const [
              DropdownMenuItem(
                  child: SizedBox(
                width: 70,
                child: Text(
                  "Hello world",
                  overflow: TextOverflow.ellipsis,
                ),
              ))
            ],
            onChanged: (value) {},
            underline: const SizedBox(),
            // selectedItemBuilder: (context) => [Container()],
          ),
        )

        // child: Row(
        //   children: [
        //     // Image.asset(
        //     //   "assets/images/profile_pic.png",
        //     //   height: 38,
        //     // ),
        //     if (!Responsive.isMobile(context))
        //       const Padding(
        //         padding: EdgeInsets.symmetric(horizontal: defaultPadding / 2),
        //         child: Text("ControllerA"),
        //       ),
        //     const Icon(Icons.keyboard_arrow_down),
        //   ],
        // ),
        );
  }
}
