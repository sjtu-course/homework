//
//  ViewController.m
//  ChatDemo
//
//  Created by bytedance on 2020/8/24.
//  Copyright © 2020 sjtu. All rights reserved.
//

#import "ViewController.h"
#import "ChatViewController.h"
#import "AppDelegate.h"

@interface ViewController ()

@property (nonatomic, weak) UITextField *roomInput;

@property (nonatomic, weak) UITextField *nameInput;

@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
    self->available_room = [[NSMutableDictionary alloc] initWithCapacity:0];
    self.view.backgroundColor = [UIColor whiteColor];
    self.title = @"Demo";
    
    UILabel *roomLabel = [[UILabel alloc] init];
    roomLabel.frame = CGRectMake(20, 150, 40, 40);
    roomLabel.text = @"房间:";
    [self.view addSubview:roomLabel];
    
    UITextField *roomInput = [[UITextField alloc] init];
    roomInput.borderStyle = UITextBorderStyleRoundedRect;
    roomInput.frame = CGRectMake(roomLabel.frame.origin.x + roomLabel.frame.size.width + 10, 150, self.view.frame.size.width - 100, 40);
    roomInput.placeholder = @"输入房间号";
    [self.view addSubview:roomInput];
    _roomInput = roomInput;
    
    UILabel *nameLabel = [[UILabel alloc] init];
    nameLabel.frame = CGRectMake(20, 250, 40, 40);
    nameLabel.text = @"用户:";
    [self.view addSubview:nameLabel];
    
    UITextField *nameInput = [[UITextField alloc] init];
    nameInput.borderStyle = UITextBorderStyleRoundedRect;
    nameInput.frame = CGRectMake(roomInput.frame.origin.x, 250, roomInput.frame.size.width, 40);
    nameInput.placeholder = @"输入用户名";
    [self.view addSubview:nameInput];
    _nameInput = nameInput;
    
    UIButton *enterButton = [UIButton buttonWithType:UIButtonTypeCustom];
    enterButton.frame = CGRectMake(50, 350, self.view.frame.size.width - 100, 45);
    enterButton.backgroundColor = [UIColor colorWithRed:0 green:0.67 blue:1.0 alpha:1.0];
    enterButton.layer.cornerRadius = 4;
    enterButton.clipsToBounds = YES;
    [enterButton setTitle:@"进入房间" forState:UIControlStateNormal];
    [enterButton setTitleColor:[UIColor whiteColor] forState:UIControlStateNormal];
    [enterButton addTarget:self action:@selector(enterButtonTapped:) forControlEvents:UIControlEventTouchUpInside];
    [self.view addSubview:enterButton];
    
    UITapGestureRecognizer *tapGesture = [[UITapGestureRecognizer alloc] initWithTarget:self action:@selector(tapRecognized)];
    [self.view addGestureRecognizer:tapGesture];
    
}

- (void)tapRecognized {
    [self.view endEditing:YES];
}

- (void)enterButtonTapped:(id)sender {
    ChatViewController *chatController;
    if([self->available_room  objectForKey:_roomInput.text]){
        chatController = [self->available_room objectForKey:_roomInput.text];
        printf("already exist\n");
        [chatController setUser:_roomInput.text user:_nameInput.text];
//      [chatController tableView:chatController->messageTable titleForHeaderInSection:@0];
    }else{
        chatController = [[ChatViewController alloc] initWithRoom:_roomInput.text user:_nameInput.text];
        [self->available_room setObject:chatController forKey:_roomInput.text];
//        for (NSString *key in self->available_room){
//            NSLog(@"key = %@",key);
//            printf("---");
//        }
        printf("create new\n");
    }
    
    
    [self.navigationController pushViewController:chatController animated:YES];
    
    [self.view endEditing:YES];
}

@end
